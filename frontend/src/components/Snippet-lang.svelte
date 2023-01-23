<script>
    export let language;
</script>

<div class="snippet-lang">
    <span>icon</span>
    <p class="snippet-lang">{language}</p>
</div>

<style>
    .snippet-lang {
        display: flex;
        position: relative;
        box-sizing: border-box;
        flex-direction: row;
        flex-wrap: nowrap;
        align-items: center;
        justify-content: center;
        /* background-color: rgb(139, 148, 8); */
        padding: 3px 10px;
    }

    .snippet-lang > * {
        display: block;
        position: relative;
        box-sizing: border-box;
        padding: 2px;
        margin: 3px;
        font-size: small;
        user-select: none;
    }
</style>
