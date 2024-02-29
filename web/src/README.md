# Vue 3 + TypeScript + Vite

## Used packages

- [Pinia](https://github.com/vuejs/pinia) (2.1.7) - state management
- [vue3-notification](https://github.com/kyvg/vue3-notification) (3.2.0) - notifications
- [vanilla-jsoneditor](https://github.com/josdejong/svelte-jsoneditor) (0.21.6) - json editor
- [tailwindcss](https://github.com/tailwindlabs/tailwindcss) (3.4.1) - css framework
- [axios](https://github.com/axios/axios) (1.6.7) - http client

----
## Develop & Debug & Test

Recommended IDE - [VS Code](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin).

1) Install [Node.js](https://nodejs.org/en/download) version 18+. 20+
2) Install dependencies:
```bash
  npm install
```
3) Start vite dev server:
```bash
  npm run dev
```

----

## Production
Build project: 
```bash
npm run build
```

----

## All Scripts

- `npm run dev` - Start Vite dev server in the current directory.
- `npm run build` - production build
- `npm run preview` - Locally preview the production build. Do not use this as a production server as it's not designed for it.

----



## Type Support For `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin) to make the TypeScript language service aware of `.vue` types.

If the standalone TypeScript plugin doesn't feel fast enough to you, Volar has also implemented a [Take Over Mode](https://github.com/johnsoncodehk/volar/discussions/471#discussioncomment-1361669) that is more performant. You can enable it by the following steps:

1. Disable the built-in TypeScript Extension
   1. Run `Extensions: Show Built-in Extensions` from VSCode's command palette
   2. Find `TypeScript and JavaScript Language Features`, right click and select `Disable (Workspace)`
2. Reload the VSCode window by running `Developer: Reload Window` from the command palette.
