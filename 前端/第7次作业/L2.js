url = "https://taskapi.chovrio.club/users/register/"
fetch(url, {
        method: "post",
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify({
            "username": "yydy",
            "password": "s7g"
        }),
    }
).then(response => response.text())
    .then(data =>console.log(data))