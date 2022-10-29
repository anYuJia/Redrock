let arr = [1, [2, 3], [4, 5, [6, 7, 8]], 9];
function flatten(arr) {
    resault = [];
    for (let i = 0; i < arr.length; i++) {
        if (Array.isArray(arr[i])) {
            resault = resault.concat(flatten(arr[i]));
        }
        else {
            resault.push(arr[i]);
        }
    }
    return resault;
}
function flatten2(arr){
    return arr.reduce((result,item)=>{
        if(Array.isArray(item)){
            return result.concat(flatten2(item))
        }
        else return result.concat(item);
    },[])
}
arr1 = flatten(arr)
arr2 = flatten2(arr)
console.log(arr1)
console.log(arr1)
//结果 [1,2,3,4,5,6,7,8,9]