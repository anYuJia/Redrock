const shallowCopy = (obj)=>{
    if(obj instanceof Array){
        let copyArr = []
        let i = 0
        for (i = 0;i<obj.length;i++) {
            copyArr[i] = obj[i]
        }
        return copyArr
    }
    else{
        let copyObj = {}
        for (let Key in obj) {
            copyObj[Key] = obj[Key]
        }
        return copyObj
    }

}
//对于对象这个属性浅拷贝拷贝后对象会影响原对象的值
//对于简单的属性不会影响原对象的值
let obj1 = {
    name:"yyy",
    age:12,
    info:{
        home:'cq',
        add:'ffff'
    }
}
let obj2 = shallowCopy(obj1)
//obj2.info.home = 'hn'
//obj2.age = 18
//console.log(obj1.age)
//console.log(obj1.info.home)

//对于数组来说，浅拷贝之后，直接的值不会影响原数组的值
let arr1 = [1,2,3,4,[5,6,7],{name:'yyy'}]
let arr2 = shallowCopy(arr1)
arr2[1]= 10
arr2[4] = [8,9,10]
arr2[5] = {name:'zzz'}
//console.log(arr1[1])
//console.log(arr1[4])
//console.log(arr1[5])

//深拷贝
const deepCopy = (obj)=>{
    if(obj instanceof Array){
        let copyArr = []
        let i = 0
        for (let o in obj) {
            if(obj[o] instanceof Array){
                copyArr[i] =deepCopy(obj[o])
            }
            if(obj[o] instanceof Object){
                copyArr[i] =deepCopy(obj[o])
            }
            else{
                copyArr[i] = obj[o]
            }
        }
        return copyArr
    }
    if(obj instanceof Object){
        let copyObj = {}
        for (let o in obj) {
            if(obj[o] instanceof Array){
                copyObj[o] = deepCopy(obj[o])
            }
            if(obj[o] instanceof Object){
                copyObj[o] = deepCopy(obj[o])
            }
            else{
                copyObj[o] = obj[o]
            }
        }
        return copyObj
    }
}
obj3 = deepCopy(obj1)
obj3.info.home = '深拷贝'
console.log(obj1.info.home)

//实现深拷贝，想了好久，用递归结合上一次的作业一层层的实现
