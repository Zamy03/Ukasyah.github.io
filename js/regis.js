import { endpointRegister } from "../js/url.js";

document.querySelector(".account-form").addEventListener("submit", async function (event) {
    event.preventDefault();

    const formData = new FormData(this);
    const data = {
        username: formData.get('username'),
        email: formData.get('email'),
        password: formData.get('password')
    };

    console.log('Form data:', data);

    try {
        const response = await fetch(endpointRegister, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });

        console.log('Response status:', response.status);
        console.log('Response headers:', response.headers);

        const responseData = await response.json();
        console.log('Response data:', responseData);

        if (!response.ok) {
            alert(responseData.message || "Registration failed.");
            return;
        }

        alert("Registration successful! Please log in.");
        window.location.replace("/login.html");

    } catch (error) {
        console.error("Error during registration:", error);
        alert("An error occurred. Please try again.");
    }
});

function getUserDetails() {
    return fetch(endpointRegister, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
    }).then((response) => response.json());

}