package template

var FRONTEND_TEMPLATES = []Item{
	{Id: 0, Title: "vue", Desc: "Generate Vue.js App Template", Command: "npm", CommandArgs: "init --y vite@latest my-vue-app -- --template vue"},
	{Id: 1, Title: "vue-ts", Desc: "Generate Vue.js App Template in TypeScript", Command: "npm", CommandArgs: "init --y vite@latest my-vue-typescript-app -- --template vue-ts"},
	{Id: 2, Title: "react", Desc: "Generate React.js App Template", Command: "npm", CommandArgs: "init --y vite@latest my-react-app -- --template react"},
	{Id: 3, Title: "react-ts", Desc: "Generate React.js App Template in TypeScript", Command: "npm", CommandArgs: "init --y vite@latest my-react-typescript-app -- --template react-ts"},
	// TODO: move to prompt
	{Id: 4, Title: "next", Desc: "Generate Next.js App Template", Command: "npx", CommandArgs: "--yes create-next-app my-next-app --eslint --src-dir --experimental-app false --use-npm --import-alias '@/*' --js"},
	{Id: 5, Title: "next-ts", Desc: "Generate Next.js App Template in TypeScript", Command: "npx", CommandArgs: "--y create-next-app my-next-typescript-app --eslint --src-dir --experimental-app false --use-npm --import-alias '@/*' --ts"},
	{Id: 6, Title: "vanilla", Desc: "Generate Vanilla.js App Template", Command: "npm", CommandArgs: "init --y vite@latest my-vanilla-app -- --template vanilla"},
	{Id: 7, Title: "vanilla-ts", Desc: "Generate Vanilla.js App Template in TypeScript", Command: "npm", CommandArgs: "init --y vite@latest my-vanilla-typescript-app -- --template vanilla-ts"},
	{Id: 8, Title: "gatsby", Desc: "Generate Gatsby App Template", Command: "npm", CommandArgs: "init --y vite@latest my-gatsby-app -- --template gatsby"},
	{Id: 9, Title: "gatsby-ts", Desc: "Generate Gatsby App Template in TypeScript", Command: "npm", CommandArgs: "init --y vite@latest my-gatsby-typescript-app -- --template gatsby-ts"},
}
