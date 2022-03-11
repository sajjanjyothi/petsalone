export const getAll = () => {
  const bearer = "Bearer " + localStorage.getItem("authtoken");

  return fetch("/api/pets", {
    credentials: 'include',
    headers: new Headers({ Authorization: bearer }),
  })
    .then((response) => {
      if (response.status === 401) window.location.href = "/login";
      else response.json();
    })
    .then((response) => response);
};
