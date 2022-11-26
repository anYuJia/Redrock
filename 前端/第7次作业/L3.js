url = "https://taskapi.chovrio.club/chatroom"
fetch(url, {
        method: "get",
        headers: {
            'Authorization':"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NCwidXNlcm5hbWUiOiJ5eWQiLCJhdmF0YXIiOiIiLCJpYXQiOjE2Njk0NDk4MDIsImV4cCI6MTY3MjA0MTgwMn0.9mNVArq9hm4g45iaN73XKT-ru0wVBGTf2OlQM-g2CZ0"
        },
    }
).then(response => response.text())
    .then(data =>console.log(data))