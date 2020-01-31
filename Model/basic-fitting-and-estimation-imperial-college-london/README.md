# idyll-post

Template for a single Idyll post.

## Installation

- Make sure you have `idyll` installed (`npm i -g idyll`).
- Clone this repo and run `npm install`.

## Hacks

idyll使用最基本的`idyll`现在的新版本会出一些问题。转而使用以下的命令
```
idyll --ssr=false
```
具体为什么我也搞不清，可见 https://gitter.im/idyll-lang/Lobby

## 作图

使用vega-lite

```
npm install --save idyll-vega-lite    

```

## Developing a post locally

Run `idyll`.

## Building a post for production

Run `idyll build`. The output will appear in the top-level `build` folder. To change the output location, change the `output` option in `package.json`.

## Deploying

Make sure your post has been built, then deploy the docs folder via any static hosting service.
