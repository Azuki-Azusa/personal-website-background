var text = "";
text += "<el-container>";
text += "<el-main>";
text += "<el-row :gutter=\"20\">";
text += "<el-col :span=\"12\">";
text += "<el-container>";
text += "<el-header>";
text += "<el-select style=\"width:100%\" v-model=\"value\" value-key=\"ID\" placeholder=\"请选择\" @change=\"optionChange\">";
text += "<el-option v-for=\"item in options\" :key=\"item.ID\" :label=\"item.Title\" :value=\"item\">";
text += "</el-option>";
text += "</el-select>";
text += "<el-input v-model=\"title\" placeholder=\"请输入题目\"></el-input>";
text += "</el-header>";
text += "<el-main>";
text += "<el-input :autosize=\"{ minRows: 28 }\" type=\"textarea\" placeholder=\"请输入内容(markdown)\" v-model=\"textarea\" @change=\"textChange\">";
text += "</el-input>";
text += "</el-main>";
text += "</el-container>"
text += "</el-col>";
text += "<el-col id=\"markcol\" :span=\"12\">";
text += "<span v-html=\"content\" v-highlight></span>";
text += "</el-col>";
text += "</el-row>";
text += "</el-main>";
text += "<el-footer style=\"text-align: right;\">";
text += "<el-button @click=\"commit\" type=\"primary\">提交</el-button>";
text += "</el-footer>";
text += "</el-container>"


const BlogList = {
    data: function () {
        return {
            options: [],
            value: '',
            title: '',
            textarea: '',
            content: ''
        }
    },
    methods: {
        optionChange: function() {
            this.title = this.value.Title
            this.textarea = this.value.Markdown
            this.content = this.value.Html
        },
        textChange: function () {
            this.content = marked(this.textarea);
        },
        commit: function () {
            var _this = this
            axios.put('/blog', {
                token: $.cookie('token'),
                id: _this.value.ID,
                title: _this.title,
                textarea: _this.textarea,
                content: _this.content
            })
                .then(function (response) {
                    alert("msg: " + response.data.msg);
                })
                .catch(function (error) {
                    alert(error)
                });
        },
    },
    created: function () {
        var _this = this
        axios.get('/bloglist?token=' + $.cookie('token'))
            .then(function (response) {
                _this.options = response.data.bloglist;
            })
            .catch(function (error) {
                alert(error)
            });
    },
    directives: {
        highlight(el) {
            let blocks = el.querySelectorAll('pre code');
            blocks.forEach((block) => {
                hljs.highlightBlock(block)
            })

            $("#markcol img").css({
                "max-width": "100%",
                "display": "block",
                "margin": "0 auto"
            })
        }
    },
    template: text
}

export default BlogList;