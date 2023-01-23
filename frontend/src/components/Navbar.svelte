<script>
</script>

<nav id="main-nav">
    <h1 class="app-title-hero">Snipux</h1>
    <button class="snippet-nav-btn"> Sort by... </button>
    <button class="snippet-nav-btn"> Add Snippet </button>
</nav>

<style>
    #main-nav {
        display: flex;
        position: sticky;
        top: 0;
        box-sizing: border-box;
        flex-direction: row;
        flex-wrap: nowrap;
        align-items: center;
        justify-content: center;
        width: 100%;
        padding: 3px 10px;
        z-index: 1;
        background-color: rgba(27, 38, 54, 0.9);
        border: solid 0px transparent;
        border-radius: 6px;
        box-shadow: 0px 1px 6px 0 rgb(27 38 54 / 80%);
    }

    #main-nav > * {
        margin: 3px;
    }

    .app-title-hero {
        flex-grow: 2;
        margin: 0;
        user-select: none;
    }

    .snippet-nav-btn {
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

    .snippet-nav-btn:hover {
        box-shadow: inset 1px 2px 3px 1px #6a27a9;
        translate: 0px 1px;
    }

    .snippet-nav-btn:active {
        box-shadow: inset 2px 3px 3px 1px #5e2495;
        background: #6a27a9;
        translate: 0px 2px;
    }
</style>
