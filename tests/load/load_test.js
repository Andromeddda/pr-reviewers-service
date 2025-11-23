import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    vus: 100, // virtual users
    duration: '30s',
    thresholds: {
        http_req_duration: ['p(99)<300'], // 99% are faster than 300s 
    },
};

export default function () {
    // Command
    let payloadTeam = JSON.stringify({
        team_name: `team-${__VU}`,
        members: [
            { user_id: `u-${__VU}-1`, username: "Alice", is_active: true },
            { user_id: `u-${__VU}-2`, username: "Bob", is_active: true },
            { user_id: `u-${__VU}-3`, username: "Carl", is_active: true },
            { user_id: `u-${__VU}-4`, username: "Daniel", is_active: true },
            { user_id: `u-${__VU}-5`, username: "Elisabeth", is_active: true },
            { user_id: `u-${__VU}-6`, username: "Forrest", is_active: true },
            { user_id: `u-${__VU}-7`, username: "Geyer", is_active: true },
            { user_id: `u-${__VU}-8`, username: "Hamlet", is_active: true },
            { user_id: `u-${__VU}-9`, username: "Ieva", is_active: true },
            { user_id: `u-${__VU}-10`, username: "Jack", is_active: true }
        ]
    });

    let headers = { 'Content-Type': 'application/json' };

    let res = http.post(
        'http://prs:8080/team/add',
        payloadTeam,
        { headers }
    );

    check(res, { 'team created or exists': r => r.status === 201 || r.status === 400 });

    // Pull request
    let prPayload = JSON.stringify({
        pull_request_id: `pr-${__VU}-${Date.now()}`,
        pull_request_name: "LoadTest PR",
        author_id: `u-${__VU}-1`,
    });

    let prRes = http.post(
        'http://prs:8080/pullRequest/create',
        prPayload,
        { headers }
    );

    check(prRes, { 'PR created': r => r.status === 201 || r.status === 409 });

    sleep(0.1);
}
