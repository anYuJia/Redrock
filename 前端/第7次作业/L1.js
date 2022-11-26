url = "https://taskapi.chovrio.club/poetry/1314"
url2 = "https://taskapi.chovrio.club/poetry"
fetch(url)
    .then(response => response.json())
    .then(data => {
        console.log('title:'+data['title'])
        fetch(url2+'?title='+data['title'])
            .then(response => response.json())
            .then(data =>console.log('content:'+data['content']))
    });