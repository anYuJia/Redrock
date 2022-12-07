function say(age,sex){
    console.log("hello,world!"+this.name+"  "+age+"  "+sex);
}
let cat = {
    name:"五仁"
}
let dog = {
    name:"来宝"
}
Function.prototype.myCall = function (obj){
    obj = obj || window;
    obj.func = this;
    let myArguments = []
    for (let i = 1;i<arguments.length;i++){
        myArguments.push('arguments['+i+']');
    }
    eval('obj.func('+myArguments+')');
    delete obj.func;
}
Function.prototype.myApply = function (obj){
        obj = obj || window;
        obj.func = this;
        let myArguments = [];
        for (let i = 0;i < arguments[1].length;i++){
            myArguments.push('arguments[1]['+i+']');
         }
        eval('obj.func('+myArguments+')');
        delete  obj.func;
}

say.myApply(cat,[1,'m'])
say.myCall(dog,2,'f')