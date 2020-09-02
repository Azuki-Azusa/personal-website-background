import BlogCreate from './BlogCreate.js'
import BlogList from './BlogList.js'
import LyricCreate from './LyricCreate.js'
import LyricList from './LyricList.js'

const NotFound = { template: '<p>Page not found</p>' }

var app = new Vue({
    el: '#app',
    components: {
        'blog-create': BlogCreate,
        'blog-list': BlogList,
        'lyric-create': LyricCreate,
        'lyric-list': LyricList,
        'not-found': NotFound
    },
    data: {
        currentTab: '1-1',
        tabs: {
            '1-1': 'blog-create',
            '1-2': 'blog-list',
            '2-1': 'lyric-create',
            '2-2': 'lyric-list',
        }
    },
    computed: {
        currentTabComponent: function () {
            return this.tabs[this.currentTab] || 'not-found'
        }
    },
    methods: {
        handleSelect(index, indexPath) {
            this.currentTab = index;
        }
    },
})