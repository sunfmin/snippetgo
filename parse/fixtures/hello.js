
export default {
    // @snippet_begin(HelloworldJS)
    components: {
        EditorContent,
        EditorMenuBubble,
        Icon,
    },

    props: ['value'],
    // @snippet_end

    data() {
        return {
            editor: new Editor({
                content: this.$props.value,
                extensions: [
                    new Blockquote(),
                    new BulletList(),
                ]
            })
        }
    },

    beforeDestroy() {
        this.editor.destroy()
    }
}
