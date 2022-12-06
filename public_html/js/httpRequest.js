function getFetchOptions(method, data) {
    const opts = {};
    opts.method = method

    if (data !== undefined) {
        opts.body = data
        opts.headers = {
            'Content-Type': 'application/json'
        }
    }

    return opts;
}

async function httpRequest(method, url, data) {
    const opts = getFetchOptions(method, data);
    console.log("before fetch")
    const resp = await fetch(url, opts);
    await resp.status
    console.log("after fetch")

    console.log("resp: ", resp)
    switch (resp.status) {
        case 200:
            console.log("hi 200")
            return handleResp(resp);
        case 204:
            console.log("hi 204")
            return null;
        case 400:
            console.log("hi 400")
            return handleResp(resp);
        default:
            throw (`error encountered, status code of ${resp.status}`)
    }
}


const handleResp = async resp => {
    const parsed = await resp.json();
    console.log("responce parsed")
    if (!!parsed.errors && parsed.errors.length > 0) {
        parsed.errors.forEach(notifyError);
        // TODO: Consider throwing here
        return null;
    }

    console.log("data after parse: ", parsed.data)
    return parsed.data;
};

const notifyError = err =>
    console.log(err, "error", 400)
