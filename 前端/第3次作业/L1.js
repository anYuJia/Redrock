function muPush(arr,value){
    arr.splice(arr.length,0,value);
    return arr;
}

//测试代码，比如向下面的arr尾部增加2
let arr = [1];
arr = muPush(arr,2)
console.log(arr)