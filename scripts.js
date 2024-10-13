function handleActionChange() {
    const action = document.getElementById("action").value;
    const formSection = document.getElementById("form-section");
    formSection.innerHTML = ""; // Clear any existing form

    if (action === "takeback") {
        formSection.innerHTML = `
            <h3>Takeback</h3>
            <label for="friend-name">Friend's Name:</label>
            <input type="text" id="friend-name" placeholder="Enter friend's name">
            <button onclick="takeback()">Takeback Item</button>
        `;
    } else if (action === "give") {
        formSection.innerHTML = `
            <h3>Give</h3>
            <label for="friend-name">Friend's Name:</label>
            <input type="text" id="friend-name" placeholder="Enter friend's name">
            <label for="item-name">Item:</label>
            <input type="text" id="item-name" placeholder="Enter item name">
            <button onclick="giveItem()">Lend Item</button>
        `;
    } else if (action === "newfriend") {
        formSection.innerHTML = `
            <h3>New Friend</h3>
            <label for="friend-name">Friend's Name:</label>
            <input type="text" id="friend-name" placeholder="Enter new friend's name">
            <button onclick="addFriend()">Add Friend</button>
        `;
    }
}

function takeback() {
    const friendName = document.getElementById("friend-name").value;
    fetch(`/api/takeback?friend=${encodeURIComponent(friendName)}`, { method: 'POST' })
        .then(response => response.json())
        .then(data => displayResult(data.message))
        .catch(error => displayResult("Error taking back item."));
}

function giveItem() {
    const friendName = document.getElementById("friend-name").value;
    const itemName = document.getElementById("item-name").value;
    fetch(`/api/give?friend=${encodeURIComponent(friendName)}&item=${encodeURIComponent(itemName)}`, { method: 'POST' })
        .then(response => response.json())
        .then(data => displayResult(data.message))
        .catch(error => displayResult("Error lending item."));
}

function addFriend() {
    const friendName = document.getElementById("friend-name").value;
    fetch(`/api/newfriend?friend=${encodeURIComponent(friendName)}`, { method: 'POST' })
        .then(response => response.json())
        .then(data => displayResult(data.message))
        .catch(error => displayResult("Error adding friend."));
}

function displayResult(message) {
    const resultSection = document.getElementById("result");
    resultSection.innerText = message;
}
