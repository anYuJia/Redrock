url = "https://taskapi.chovrio.club/"
fetch(url)
    .then(response => response.json())
    .then(data => console.log(data));