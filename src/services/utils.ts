export const headers = {
  'Content-Type': 'application/json',
  Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};
