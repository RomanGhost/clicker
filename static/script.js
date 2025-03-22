let ws;
const clickButton = document.getElementById('click-button');
const validClicksDisplay = document.getElementById('valid-clicks');
const allClicksDisplay = document.getElementById('all-clicks');
const messageArea = document.getElementById('message-area');
const userLoginDisplay = document.getElementById('user-login');
const userLeagueDisplay = document.getElementById('user-league');
const updateList = document.getElementById('update-list');
const clickValueDisplay = document.getElementById('click-value-display');

const loginForm = document.getElementById('login-form');
const signupForm = document.getElementById('signup-form');
const toggleAuthFormButton = document.getElementById('toggle-auth-form');
const loginButton = document.getElementById('login-button');
const signupButton = document.getElementById('signup-button');
const loginUsernameInput = document.getElementById('login-username');
const loginPasswordInput = document.getElementById('login-password');
const signupUsernameInput = document.getElementById('signup-username');
const signupPasswordInput = document.getElementById('signup-password');
const signupAcceptPasswordInput = document.getElementById('signup-accept-password');

let userLogin = "";
let clicksBatch = [];
let clickCounter = 0;
let currentClickValue = 1;

function displayMessage(message, type = 'info') {
    const messageElement = document.createElement('div');
    messageElement.textContent = message;
    messageElement.classList.add(type === 'error' ? 'error-message' : 'success-message');
    messageArea.appendChild(messageElement);
    messageArea.scrollTop = messageArea.scrollHeight;
}

function setupWebSocket() {
    if (ws) ws.close();
    ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => displayMessage("WebSocket connected", 'success');
    ws.onmessage = event => {
        const message = JSON.parse(event.data);
        console.log("Received:", message);
        if (message.typeMessage === 'user') {
            validClicksDisplay.textContent = message.data.valid_clicks;
            allClicksDisplay.textContent = message.data.all_clicks;
            userLoginDisplay.textContent = message.data.login || userLogin;
            userLeagueDisplay.textContent = message.data.league || "Rookie";
            clickValueDisplay.textContent = message.data.click_value || 1;
        }
    };
    ws.onclose = () => displayMessage("WebSocket closed", 'error');
}

clickButton.onclick = function() {
    clickCounter++;
    clicksBatch.push({ click_value: currentClickValue, click_time: Math.floor(Date.now() / 1000) });
    if (clickCounter >= 20) sendClickBatch();
};

function sendClickBatch() {
    if (clicksBatch.length === 0) return;
    ws.send(JSON.stringify({ TypeMessage: "click_batch", Data: { clicks_info: clicksBatch } }));
    clicksBatch = [];
}

toggleAuthFormButton.onclick = function() {
    loginForm.style.display = loginForm.style.display === 'none' ? 'block' : 'none';
    signupForm.style.display = signupForm.style.display === 'none' ? 'block' : 'none';
};

loginButton.onclick = function() {
    fetch('/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ login: loginUsernameInput.value, password: loginPasswordInput.value }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.message === "logged in successfully") {
            displayMessage("Logged in", 'success');
            userLogin = loginUsernameInput.value;
            setupWebSocket();
            clickButton.disabled = false;
        }
    });
};
