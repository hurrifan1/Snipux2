<script>
    import SnippetLang from "./Snippet-lang.svelte";
    import Toast from "./Toast.svelte";
    import { toast } from '../data/stores.js';

    export let name;
    export let language;
    export let description;
    export let lastModifiedGMT;
    export let code;

    function copyCode() {
        console.log(">>>> Live from the `copyCode()` function!!");
        navigator.clipboard.writeText(code);
        $toast = "Copied!";
        console.log(">>>> Toast is now:", $toast);
    }
</script>

<div class="snippet-body">
    <div class="snippet-bar">
        <button class="snippet-action-btn" on:click={copyCode}>Copy</button>
        <button class="snippet-action-btn">Edit</button>
        <button class="snippet-action-btn">Delete</button>
        <h3 class="snippet-title">{name}</h3>
        <SnippetLang {language} />
        <p class="snippet-last-mod-date">
            {new Date(lastModifiedGMT * 1000).toLocaleDateString()}
        </p>
        <!--  new Date(1669684930000).toLocaleDateString()  -->
    </div>
    <p class="snippet-desc">{description}</p>
    <div class="snippet-preview-area">
        <code class="snippet-preview">
            {code}
        </code>
    </div>
</div>

<style>
    .snippet-body {
        display: block;
        position: relative;
        box-sizing: border-box;
        /* background-color: aquamarine; */
        background-color: rgba(255, 255, 255, 0.1);
        margin: 30px 0;
        padding: 10px 2%;
        border: solid 0px transparent;
        border-radius: 10px;
        box-shadow: 0 0 7px 0px rgb(27 38 54 / 90%);

        transition: all 0.1s ease-in-out;
    }

    .snippet-body:hover {
        box-shadow: 0 0 7px 7px rgb(27 38 54 / 90%);
    }

    .snippet-body:nth-child(1) {
        margin-top: 15px;
        margin-bottom: 30px;
    }

    .snippet-bar {
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

    .snippet-bar > * {
        display: block;
        position: relative;
        box-sizing: border-box;
        padding: 2px;
        margin: 3px;
    }

    .snippet-action-btn {
        background: blueviolet;
        color: white;
        border: solid 1px #6a27a9;
        border-radius: 4px;
        padding: 7px;
        cursor: pointer;
        user-select: none;
        box-shadow: 1px 3px 3px 0 #6a27a9;

        transition: all 0.1s ease-in-out;
    }

    .snippet-action-btn:hover {
        box-shadow: inset 1px 2px 3px 1px #6a27a9;
        translate: 0px 1px;
    }

    .snippet-action-btn:active {
        box-shadow: inset 2px 3px 3px 1px #5e2495;
        background: #6a27a9;
        translate: 0px 2px;
    }

    .snippet-title {
        flex-grow: 1;
        user-select: none;
    }

    .snippet-last-mod-date {
        /* display: block; */
        position: relative;
        box-sizing: border-box;
        text-align: center;
        font-size: small;
        user-select: none;
    }

    .snippet-desc {
        display: block;
        position: relative;
        box-sizing: border-box;
        margin: 10px 3px;
        text-align: left;
        font-size: small;
        font-style: italic;
        user-select: none;
    }

    .snippet-preview-area {
        background-color: rgb(0 0 0 / 20%);
        margin: 10px 0;
        padding: 10px;
        text-align: left;
        border: solid 0px transparent;
        border-radius: 6px;
    }

    .snippet-preview {
        color: white;
        white-space: pre-wrap;
    }
</style>
