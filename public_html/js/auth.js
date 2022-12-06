function loginSubmit(evt) {
    evt.preventDefault()
    var email = document.getElementById("emailSubmit")
    var password = document.getElementById("passwordSubmit")
    const url = "/api/login"

    fetch(url, {
        method: 'POST',
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            email: email.value,
            password: password.value,
        }),
    })
        .then(response => response.json())
        .then(json => console.log("response: " + json));
}

/*
function loginSubmit(evt) {
    evt.preventDefault()
    const url = "/api/login"
    data = JSON.stringify({
        email: document.getElementById("emailSubmit").value,
        password: document.getElementById("passwordSubmit").value
    })

    resp = httpRequest("POST", url, data)

}
*/

/*
function signUp(email, password, repeatPassword) {
    if (password != repeatPassword) {
        return
    }
    headers.append('Authorization', 'Basic' + base64.encode(email + ":" + password));

    fetch(url, {
        method: 'GET',
        headers: headers,
        //credentials: 'user:passwd'
    })
        .then(response => response.json())
        .then(json => console.log(json));
};
*/
