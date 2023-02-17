package template

var BACKEND_TEMPLATES = []Item{
	{Title: "express", Desc: "Generate Express.js App Template", Command: "npx", CommandArgs: "--yes express-generator my-express-app"},
	{Title: "koa", Desc: "Generate Koa.js App Template", Command: "npx", CommandArgs: "--yes create-koa-application my-koa-app"},
}
