import { endpointGetTentara, endpointPostTentara, endpointPutTentara, endpointDeleteTentara } from "../js/url.js";

function fetchTentara() {
    const token = localStorage.getItem("token");
    fetch(endpointGetTentara, {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })
    .then(response => response.json())
    .then(data => {
        const tbody = document.getElementById('tentara-tbody');
        tbody.innerHTML = '';
        data.forEach(tentara => {
            const tr = document.createElement('tr');
            tr.innerHTML = `
                <td contenteditable="true" data-field="nama_lengkap">${tentara.nama_lengkap}</td>
                <td contenteditable="true" data-field="pangkat">${tentara.pangkat}</td>
                <td contenteditable="true" data-field="tanggal_lahir">${tentara.tanggal_lahir}</td>
                <td contenteditable="true" data-field="tanggal_wafat">${tentara.tanggal_wafat}</td>
                <td contenteditable="true" data-field="nomor_identitas">${tentara.nomor_identitas}</td>
                <td>
                    <button onclick="editTentara(this, ${tentara.tentara_id})">Save</button>
                    <button onclick="deleteTentara(${tentara.tentara_id})">Delete</button>
                </td>
            `;
            tbody.appendChild(tr);
        });
    })
    .catch(error => console.error('Error fetching data:', error));
}

function showTentaraForm() {
    document.getElementById('tentara-form').style.display = 'block';
}

function addTentara() {
    const form = document.getElementById('new-tentara-form');
    const newTentara = {
        nama_lengkap: form.nama_lengkap.value,
        pangkat: form.pangkat.value,
        tanggal_lahir: form.tanggal_lahir.value,
        tanggal_wafat: form.tanggal_wafat.value,
        nomor_identitas: form.nomor_identitas.value
    };

    const token = localStorage.getItem("token");

    fetch(endpointPostTentara, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(newTentara)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
    })
    .then(data => {
        console.log('Tentara added:', data);
        fetchTentara(); // Refresh the tentara list
        form.reset(); // Clear the form
        document.getElementById('tentara-form').style.display = 'none'; // Hide the form
    })
    .catch(error => console.error('Error adding tentara:', error));
}

function editTentara(button, id) {
    const tr = button.closest('tr');
    const updatedTentara = {};
    tr.querySelectorAll('[contenteditable="true"]').forEach(td => {
        const field = td.getAttribute('data-field');
        updatedTentara[field] = td.textContent;
    });

    const token = localStorage.getItem("token");
    const url = endpointPutTentara.replace("{id}", id);

    fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(updatedTentara)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
    })
    .then(data => {
        console.log('Tentara updated:', data);
        fetchTentara(); // Refresh the tentara list
    })
    .catch(error => console.error('Error updating tentara:', error));
}

function deleteTentara(tentara_id) {
    const token = localStorage.getItem("token");
    const url = endpointDeleteTentara.replace("{id}", tentara_id);

    fetch(url, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
    })
    .then(data => {
        console.log('Tentara deleted:', data);
        fetchTentara(); // Refresh the tentara list
    })
    .catch(error => console.error('Error deleting tentara:', error));
}

// Panggil fungsi fetchTentara saat halaman dimuat
document.addEventListener('DOMContentLoaded', fetchTentara);

// Make functions available in the global scope
window.editTentara = editTentara;
window.deleteTentara = deleteTentara;

// Attach event listeners
document.getElementById('show-tentara-form').addEventListener('click', showTentaraForm);
document.getElementById('submit-tentara').addEventListener('click', addTentara);
