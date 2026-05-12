import http from 'k6/http';
import { check } from 'k6';

function uuidv4() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        let r = Math.random() * 16 | 0, v = c === 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
}

const targetWallets = [
    '550e8400-e29b-41d4-a716-446655440000',
    '660e8400-e29b-41d4-a716-446655441111'
];

export const options = {
    scenarios: {
        wallet_load_test: {
            executor: 'constant-arrival-rate',
            rate: 1000,
            timeUnit: '1s',
            duration: '30s',
            preAllocatedVUs: 200,
            maxVUs: 1000,
        },
    },
    thresholds: {
        'http_req_failed{status:500}': ['rate<0.01'], 
        'http_req_failed{status:501}': ['rate<0.01'],
        'http_req_failed{status:502}': ['rate<0.01'],
        'http_req_failed{status:503}': ['rate<0.01'],
        'http_req_failed{status:504}': ['rate<0.01'],
        'http_req_failed{status:524}': ['rate<0.01'],
        http_req_duration: ['p(95)<500'],
    },
};

export default function () {
    const url = 'http://localhost:8080/api/v1/wallet';
    


    const isTarget = Math.random() < 0.7;
    const currentWalletId = isTarget 
        ? targetWallets[Math.floor(Math.random() * targetWallets.length)] 
        : uuidv4();


    const operation = Math.random() < 0.5 ? 'DEPOSIT' : 'WITHDRAW';
    

    const amount = Math.floor(Math.random() * 1000) + 1;

    const payload = JSON.stringify({
        valletId: currentWalletId,
        operationType: operation,
        amount: amount,
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const response = http.post(url, payload, params);

    check(response, {
        'status is 200': (r) => r.status === 200,


        'is not 500': (r) => r.status !== 500,
    });
}
