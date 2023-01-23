import { writable } from 'svelte/store';
import { main } from '../../wailsjs/go/models';

import { GetCurrentSnippets } from "../../wailsjs/go/main/App.js";

export const modal = writable(null);
export const toast = writable(null);

// let hold = new main.Snippet();
// hold.code = "ggg";
// hold.description = "sss";
// hold.language = "JavaScript";
// hold.lastModifiedGMT = 1669790139;
// hold.name = "oyewoinowc";

// export let snippets = writable([hold]);

export let snippets = writable([]);
GetCurrentSnippets().then(r => snippets.update((z) => z = r));
