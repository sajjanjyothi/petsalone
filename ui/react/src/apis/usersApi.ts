export const authenticate = (username: string, password: string) => {
    return fetch('/api/users/login', {
        method: 'post',
        headers: new Headers({'content-type': 'application/json'}),
        body: JSON.stringify({username, password})
      }).
      then(function(response) {
        return response.json();
      }).then(function(data) {
        localStorage.setItem("authtoken", data.token);
        localStorage.setItem("loggedin","true")
        return data;
      });
}