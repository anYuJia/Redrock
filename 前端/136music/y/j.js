let step = (time) => {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve("在红岩学技术太有趣了！");
      }, time);
    });
  };
  async function demo() {
    let word1 = await step(500);
    console.log(word1); //在红岩学技术太有趣了！
    let word2 = await step(1000);
    console.log(word2);//在红岩学技术太有趣了！
    let word3 = await step(1500);
    return "函数执行完成";
  }
  demo().then((res) => {
    console.log(res);//函数执行完成
  });