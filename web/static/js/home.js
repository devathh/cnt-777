let wallet = document.getElementById("wallet");

function readLS() {
    let balance = Number(localStorage.getItem("balance"));
    return balance
}

function writeLS(balance) {
    localStorage.setItem("balance", balance);
}

function changeBalance(balance) {
    wallet.value = balance;
}

function loadBalance() {
    let balance = readLS()
    if (balance == 0) {
        writeLS(1000000)
        changeBalance(1000000)
    } else {
        changeBalance(balance)
    }
}
loadBalance();

wallet.addEventListener("blur", () => {
    let value = Number(wallet.value);
    if (value < 0) {
        writeLS(0);
        wallet.value = 0;
    } else {
        writeLS(value);
        wallet.value = value;
    }
});

let blackjackBtn = document.getElementById("blackjack-btn");
blackjackBtn.addEventListener("click", () => {
    console.log("ok")
    window.location.replace("https://www.youtube.com/watch?v=xvFZjo5PgG0")
});

let userAgreementBtn = document.getElementById("user-agreement-btn");
userAgreementBtn.addEventListener("click", () => {
    document.getElementById("user-agreement").classList.remove("hidden");
});

let acceptBtn = document.getElementById("accept-btn");
acceptBtn.addEventListener("click", () => {
    document.getElementById("user-agreement").classList.add("hidden");  
});