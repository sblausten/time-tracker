import config from '../application'


export const handleErrors = (response) => {
    if (response.status === 500) {
        const errorMessage = "Failed to save session. Please try again.";
        console.log(`Error saving session: ${response.status} - ${response.url}`);
        throw Error(errorMessage);
    }
    return response;
};

const postData = (url, data) => {
    return fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
};

const saveSession = (userId, start, end, sessionName) => {
    const SERVER_URL = process.env.SERVER_URL;
    const url = `${SERVER_URL || config.localServer}/v1/users/${userId}/session`;

    const body = {
        start: start.toISOString(),
        end: end.toISOString(),
        duration: end - start,
        name: sessionName
    };

    return postData(url, body)
};

const getSessions = (userId) => {
    const SERVER_URL = process.env.SERVER_URL;
    const url = `${SERVER_URL || config.localServer}/v1/users/${userId}/sessions`;

    return fetch(url);
};

export const client = { saveSession, getSessions };
