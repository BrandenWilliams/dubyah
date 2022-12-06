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
    const resp = await fetch(url, opts);

    switch (resp.status) {
        case 200:
            return handleResp(resp);
        case 204:
            return null;
        case 400:
            return handleResp(resp);
        default:
            throw (`error encountered, status code of ${resp.status}`)
    }
}

const handleResp = async resp => {
    const parsed = await resp.json();
    if (!!parsed.errors && parsed.errors.length > 0) {
        parsed.errors.forEach(notifyError);
        // TODO: Consider throwing here
        return null;
    }

    return parsed.data;
};

const notifyError = err =>
    console.log(err, "error", 400)
