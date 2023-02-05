const setBanner = function(){
    let url = "http://ananqiexiafan.icu:3000/banner";
    fetch(url,{
        method :"get",
    }).then(response =>response.json())
    .then(data => getBanner(data))
}
//const bannerList = document.querySelectorAll('#zi img');
const getBanner = function(data){
    for (let i = 0 ; i < data['banners'].length;i++){
      console.log(data['banners'][i]['imageUrl'])
    }
}
setBanner();