<script>
    import { username, verified } from '$lib/stores/username.js';
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';

    let tab = 'login';
    let checkValue = 'Check';

    let temp_username = '';
    $: {
        temp_username = temp_username.toLowerCase();
        checkValue = 'Check';
        $verified = false;
    }


    async function check() {
        await username.verify(temp_username);
        if ($verified) {
            checkValue = 'Verified';
        } else {
            checkValue = 'Invalid';
        }
    }

    const loginSwitch = () => {
        tab = 'login';
        checkValue = 'Check';
        $verified = false;
    }

    const loginSubmit = async (event) => {
        event.preventDefault();
        await username.verify(temp_username);
        goto('/chat');
    }

    const registerSwitch = () => {
        tab = 'register';
        checkValue = 'Check';
        $verified = false;
    }

    async function registerSubmit(event) {
        event.preventDefault();
        username.register(temp_username);
        username.verify(temp_username);
        goto('/chat');
    }

    const anonymousSwitch = () => {
        tab = 'anonymous';
    }

    async function anonymousSubmit(event) {
        event.preventDefault();
        await username.setDefault();
        goto('/chat');
    }

    onMount(async () => {
        console.log($username);
        if ($username === '') {
            await username.setDefault();
        }
        if (username.verify($username) && $username !== await username.getDefault()) {
            goto('/chat');
        }
        $verified = false;
    });
</script>

<!-- Centered with tabs login and register on top with two forms that submit to login and register functions that dont yet exist -->
<div class="flex flex-col items-center justify-center h-screen">
    <div role="tablist" class="tabs tabs-boxed">
        {#if tab === 'login'}
            <button role="tab" class="tab tab-active" on:click={loginSwitch}>Login</button>
            <button role="tab" class="tab" on:click={registerSwitch}>Register</button>
            <button role="tab" class="tab" on:click={anonymousSwitch}>Anonymous</button>
        {:else if tab === 'register'}
            <button role="tab" class="tab" on:click={loginSwitch}>Login</button>
            <button role="tab" class="tab tab-active" on:click={registerSwitch}>Register</button>
            <button role="tab" class="tab" on:click={anonymousSwitch}>Anonymous</button>
        {:else if tab === 'anonymous'}
            <button role="tab" class="tab" on:click={loginSwitch}>Login</button>
            <button role="tab" class="tab" on:click={registerSwitch}>Register</button>
            <button role="tab" class="tab tab-active" on:click={anonymousSwitch}>Anonymous</button>
        {/if}
    </div>

    {#if tab === 'login'}
        <div class="card shadow-xl flex justify-center items-center">
            <form class="box p-6 mt-4 flex flex-col items-center">
                <label for="username" class="label">Enter your username below</label>

                <div class="join">
                    <input type="text" id="username" name="username" class="input bg-base-200 join-item" placeholder="Username" bind:value={temp_username} required>
                    <button 
                        id="checkButton"
                        class="btn join-item {$verified ? 'btn-success' : 'btn-primary'}"
                        on:click={check}
                        disabled={!temp_username || $verified}
                    > { checkValue } </button>
                </div>
                <p class="text-center mt-2 text-xs">If you don't have an account, click <a on:click={registerSwitch} href="#" class="text-primary">here</a> to register</p>

                <button type="submit" class="btn btn-primary mt-4" disabled={!$verified || !temp_username}>Login</button>
            </form>
        </div>
    {:else if tab === 'register'}
        <div class="card shadow-xl flex justify-center items-center">
            <form class="box p-6 mt-4 flex flex-col items-center">
                <label for="username" class="label">Enter your username below</label>

                <div class="join">
                    <input type="text" id="username" name="username" class="input bg-base-200 join-item" placeholder="Username" bind:value={temp_username} required>
                    <button 
                        id="checkButton"
                        class="btn join-item {$verified ? 'btn-success' : 'btn-primary'}"
                        on:click={check}
                        disabled={!temp_username}
                    > { checkValue } </button>
                </div>

                <p class="text-center mt-2 text-xs">If you already have an account, click <a on:click={loginSwitch} href="#" class="text-primary">here</a> to login</p>

                <button type="submit" class="btn btn-primary mt-4" disabled={$verified || !temp_username}>Register</button>
            </form>
        </div>
    {:else if tab === 'anonymous'}
        <div class="card shadow-xl flex justify-center items-center">
            <form class="box p-6 mt-4 flex flex-col items-center" on:submit={anonymousSubmit}>
                <p class="text-center text-xs">You are about to enter as an anonymous user. Instead of a username, <br>
                                               your IP address will be used to identify you. We recommend that you <br>
                                               <a href="#" on:click={registerSwitch} class="text-primary">create an account</a> for the best experience</p>
                <p class="text-center mt-2 text-xs">If you already have an account, click <a on:click={loginSwitch} href="#" class="text-primary">here</a> to login</p>

                <button type="submit" class="btn btn-primary mt-4">Continue as Anonymous</button>
            </form>
        </div>
    {/if}

</div>

