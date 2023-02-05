const getBangDan = function(){
    let url = "http://ananqiexiafan.icu:3000/toplist/detail"
    fetch(url,{
        method : "get",
    }).then(response => response.json())
    .then(data => setBangDan(data));
}
const setBangDan = function(data){
    //title = document.querySelectorAll('.rank-list-title');
    for(let i = 0 ; i < 4 ; i++){
        console.log(data['list'][i]['name'])
        //title[i].innerHTML = data['list'][i]['name'];
        //body = document.querySelectorAll('.rank-list-body-name')
        for(let j = 0 ; j < data['list'][i]['tracks'].length ; j++){
            //body[i*3+j].innerHTML = data['list'][i]['tracks'][j];
            console.log(data['list'][i]['tracks'][j]['first'],data['list'][i]['tracks'][j]['second']);
        }
        //console.log(data['list'][i]['name'],data['list'][i]['tracks']);
    }

}
getBangDan();