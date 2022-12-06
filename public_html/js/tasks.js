// GetTask from tasklists api
async function getTask(url) {
    // Default options are marked with *
    const response = await fetch(url, {
        method: 'GET',
        credentials: 'include',
    });

    return response.json(); // parses JSON response into native JavaScript objects
}

getTask('/api/tasklists')
    .then((data) => {
        console.log(data); // JSON data parsed by `data.json()` call
    });
