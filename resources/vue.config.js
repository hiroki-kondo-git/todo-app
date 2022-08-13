// CORS対策
module.exports = {
	devServer: {
		proxy: {
			"^/api*": {
				target: "http://localhost:8082",
				pathRewrite: { "^/api": "" },
				changeOrigin: true,
				logLevel: "debug",
			}
		}
	}
};