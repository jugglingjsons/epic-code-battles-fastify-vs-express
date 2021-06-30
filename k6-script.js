import http from 'k6/http';
import { check } from 'k6';

export let options = {
  vus: 40,
  duration: '5s',
};

export default function () {
  let res = http.get('http://localhost:3000');
  check(res, {
    'is status 200': (r) => r.status === 200,
    'body size is 1176 bytes': (r) => r.body.length == 1176,
  });
}