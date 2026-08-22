// 服务入口(独立文件,避免测试时监听端口)
const { app } = require("./app");

const port = process.env.PORT || 8080;
app.listen(port, () => console.log(`listening on :${port}`));
