package template

var BACKEND_TEMPLATES = []Item{
	{Id: 0, Title: "express", Desc: "Generate Express.js App Template", Command: "npx", CommandArgs: "--yes express-generator my-express-app"},
	{Id: 1, Title: "koa", Desc: "Generate Koa.js App Template", Command: "npx", CommandArgs: "--yes create-koa-application my-koa-app"},
}
