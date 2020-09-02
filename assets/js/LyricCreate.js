var text = "";
text += "<el-container>";
text += "<el-main>";
text += "<el-row :gutter=\"20\">";
text += "<el-col :span=\"12\">";
text += "<el-container>";
text += "<el-header>";
text += "<el-input v-model=\"title\" placeholder=\"请输入题目\"></el-input>";
text += "</el-header>";
text += "<el-main>";
text += "<el-input :autosize=\"{ minRows: 28 }\" type=\"textarea\" placeholder=\"请输入内容(markdown)\" v-model=\"textarea\" @change=\"textChange\">";
text += "</el-input>";
text += "</el-main>";
text += "</el-container>"
text += "</el-col>";
text += "<el-col id=\"markcol\" :span=\"12\">";
text += "<span v-html=\"textarea\"></span>";
text += "</el-col>";
text += "</el-row>";
text += "</el-main>";
text += "<el-footer style=\"text-align: right;\">";
text += "<el-button @click=\"commit\" type=\"primary\">提交</el-button>";
text += "</el-footer>";
text += "</el-container>"


const BlogCreate = {
	data: function () {
		return {
			title: '',
			textarea: '',
		}
	},
	methods: {
		textChange: function () {
			this.content = marked(this.textarea);
		},
		commit: function () {
			var _this = this
			axios.post('/lyric', {
				token: $.cookie('token'),
				title: _this.title,
				textarea: _this.textarea
			})
				.then(function (response) {
					alert("msg: " + response.data.msg);
				})
				.catch(function (error) {
					alert(error)
				});
		},
	},
	template: text
}

export default BlogCreate;