<script>
    import { toast } from "../data/stores.js";

    // =========================================================================
    //  EXAMPLE MODAL STUFF FROM  -->  https://svelte.dev/examples/modal
    // =========================================================================
    import { createEventDispatcher, onDestroy } from 'svelte';

	const dispatch = createEventDispatcher();
	const close = () => dispatch('close');

	let modal;

	const handle_keydown = e => {
		if (e.key === 'Escape') {
			close();
			return;
		}

		if (e.key === 'Tab') {
			// trap focus
			const nodes = modal.querySelectorAll('*');
			const tabbable = Array.from(nodes).filter(n => n.tabIndex >= 0);

			let index = tabbable.indexOf(document.activeElement);
			if (index === -1 && e.shiftKey) index = 0;

			index += tabbable.length + (e.shiftKey ? -1 : 1);
			index %= tabbable.length;

			tabbable[index].focus();
			e.preventDefault();
		}
	};

	const previously_focused = typeof document !== 'undefined' && document.activeElement;

	if (previously_focused) {
		onDestroy(() => {
			previously_focused;
		});
	}

    // =========================================================================
    //  MY FUNCTIONS
    // =========================================================================

    function yesOption() {
        console.log(">>>> Do something with 'yes'");


    }

    function noOption() {
        console.log(">>>> Do something with 'no'");


    }
</script>

<div class="modal-full">
    <div class="modal-overlay">
        <div class="modal-actual">
            <h3 class="modal-title">Delete this snippet?</h3>
            <div class="modal-option-btn-container">
                <button class="modal-option-btn yes-btn" on:click={yesOption}>Yes</button>
                <button class="modal-option-btn no-btn" on:click={noOption}>No</button>
            </div>
        </div>
    </div>
</div>

<style>
    .modal-full {
        display: block;
        position: relative;
    }

    .modal-overlay {
        position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0,0,0,0.3);
    }

    .modal-actual {
        position: absolute;
		left: 50%;
		top: 50%;
		width: calc(100vw - 4em);
		max-width: 32em;
		max-height: calc(100vh - 4em);
		overflow: auto;
		transform: translate(-50%,-50%);
		padding: 1em;
		border-radius: 0.2em;
		background: white;
    }

    .modal-title {
        display: block;
        position: relative;
        user-select: none;
    }

    .modal-option-btn-container {
        display: block;
        position: relative;
    }

    .modal-option-btn {
        color: white;
        user-select: none;
    }

    .yes-btn {
        background-color: green;
    }

    .no-btn {
        background-color: darkred;
    }
</style>
