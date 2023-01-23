<!-- Modified from this excellent MDN guide ==> https://developer.mozilla.org/en-US/docs/Learn/Tools_and_testing/Client-side_JavaScript_frameworks/Svelte_stores#creating_the_alert_component -->
<script>
    // export let message;

    import { toast } from "../data/stores.js";
    import { onDestroy } from "svelte";

    import { tweened } from "svelte/motion";
    import { linear } from "svelte/easing";
    import { fade } from "svelte/transition";

    let toastTimer = 4000;
    let toastTimeout;

    let progress = tweened(1, {
        duration: toastTimer,
        easing: linear,
    });

    // progress.set(0);
    // setTimeout(() => (toastContent = ""), toastTimer);

    // const unsub = toast.subscribe((value) => (toastContent = value));
    // onDestroy(unsub);

    const onMessageChange = (message) => {
        clearTimeout(toastTimeout);
        progress.set(1);
        // progress = tweened(1, {
        //     duration: toastTimer,
        //     easing: linear,
        // });
        if ($toast) {
            // progress.set(1);
            progress.set(0);
            toastTimeout = setTimeout(() => ($toast = null), toastTimer);
        }
    };
    $: onMessageChange($toast); // whenever the alert store changes run onMessageChange

    onDestroy(() => {
        console.log("#### Toast destroyed!");
        clearTimeout(toastTimeout);
    }); // make sure we clean-up the timeout
</script>

{#if $toast}
    <div class="toast-body" on:click={() => ($toast = null)} transition:fade>
        <h3 class="toast-message">{$toast}</h3>
        <progress value={$progress} class="countdown-bar" />
    </div>
{/if}

<!-- {#if toastContent}
<div on:click={() => toastContent = ''}>
  <p>{ toastContent }</p>
</div>
{/if} -->
<style>
    .toast-body {
        position: fixed;
        cursor: pointer;
        margin-right: 1.5rem;
        margin-left: 1.5rem;
        margin-top: 1rem;
        right: 0;
        display: flex;
        align-items: center;
        border-radius: 0.2rem;
        background-color: #565656;
        color: #fff;
        font-weight: 700;
        padding: 0.5rem 1.4rem;
        font-size: 1.5rem;
        z-index: 100;
        opacity: 95%;
        width: 20%;
        flex-direction: column;

        transition: all 1s ease-in-out;
    }

    /* .toast-body.enabled {
        display: block;
        bottom: 10%;

        transition: all 1s ease-in-out;
    } */

    .toast-message {
        color: white;
    }

    .countdown-bar {
        width: 100%;
    }
</style>
