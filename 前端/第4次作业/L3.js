equal = (obj1,obj2)=>{
    let flag = 1
    for (let obj1Key in obj1) {
        if(obj1[obj1Key] !== obj2[obj1Key]){
            flag = 0
            break
        }
    }
    if(!flag){
        console.log("该两个对象不相等")
    }
    else console.log("两个对象相等")
}

boj1 = {
    name:"qqq",
    age: 12
}
boj2 = {
    name:"qqq",
    age: 12
}
equal(boj1,boj2)