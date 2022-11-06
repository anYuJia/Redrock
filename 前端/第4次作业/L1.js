const fn1 = n =>(n==1)?1:n+fn1(n-1)
const fn2 = n =>{
    let result = 0
    for (let i = 1; i <=n ; i++) {
        result += i
    }
    return result
}
console.log(fn1(3))
console.log(fn2(5))