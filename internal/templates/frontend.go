package template

var FRONTEND_TEMPLATES = []Item{
	{Title: "vue", Desc: "Generate Vue.js App Template", Command: "npm", CommandArgs: "init vite@latest my-vue-app -- --template vue"},
	{Title: "vue-ts", Desc: "Generate Vue.js App Template in TypeScript", Command: "npm", CommandArgs: "init vite@latest my-vue-typescript-app -- --template vue-ts"},
	{Title: "react", Desc: "Generate React.js App Template", Command: "npm", CommandArgs: "init vite@latest my-react-app -- --template react"},
	{Title: "react-ts", Desc: "Generate React.js App Template in TypeScript", Command: "npm", CommandArgs: "init vite@latest my-react-typescript-app -- --template react-ts"},
	{Title: "next", Desc: "Generate Next.js App Template", Command: "npm", CommandArgs: "init vite@latest my-next-app -- --template next"},
	{Title: "next-ts", Desc: "Generate Next.js App Template in TypeScript", Command: "npm", CommandArgs: "init vite@latest my-next-typescript-app -- --template next-ts"},
	{Title: "vanilla", Desc: "Generate Vanilla.js App Template", Command: "npm", CommandArgs: "init vite@latest my-vanilla-app -- --template vanilla"},
	{Title: "vanilla-ts", Desc: "Generate Vanilla.js App Template in TypeScript", Command: "npm", CommandArgs: "init vite@latest my-vanilla-typescript-app -- --template vanilla-ts"},
	{Title: "gatsby", Desc: "Generate Gatsby App Template", Command: "npm", CommandArgs: "init vite@latest my-gatsby-app -- --template gatsby"},
	{Title: "gatsby-ts", Desc: "Generate Gatsby App Template in TypeScript", Command: "npm", CommandArgs: "init vite@latest my-gatsby-typescript-app -- --template gatsby-ts"},
}
