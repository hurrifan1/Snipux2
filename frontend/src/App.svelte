<script>
  // @ts-nocheck
  import { snippets } from "./data/stores.js";

  // import { GetCurrentSettings } from "../wailsjs/go/main/App.js";
  // import { GetCurrentSnippets } from "../wailsjs/go/main/App.js";

  import Navbar from "./components/Navbar.svelte";
  import SnippetHome from "./components/Snippet-home.svelte";
  import Toast from "./components/Toast.svelte";

  // let showModal = false;
</script>

<main id="main">
  <Navbar />
  <Toast />
  <section class="snippet-container">
    {#await $snippets}
      <p>Loading snippets...</p>
    {:then value}
      {#each value as snip}
        <SnippetHome
          name={snip.name}
          language={snip.language}
          description={snip.description}
          lastModifiedGMT={snip.lastModifiedGMT}
          code={snip.code}
        />
      {/each}
    {:catch error}
      <p>Something went wrong: {error.message}</p>
    {/await}
  </section>

  <!-- {#if toast}
    <Toast message="Fart ass bitch" />
  {/if} -->
</main>

<style>
  #main {
    margin: 12px 7vw;
  }
</style>
