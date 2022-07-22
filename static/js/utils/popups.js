function buildErrorPopup(text) {
    Swal.fire({
        icon: 'error',
        title: 'Oops...',
        text: text.charAt(0).toUpperCase() + text.slice(1),
    });
}

function buildSuccessPopup(text) {
    Swal.fire({
        icon: 'success',
        showConfirmButton: false,
        text: text.charAt(0).toUpperCase() + text.slice(1),
    });
}