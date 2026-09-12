<?php

use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Http\Request;

function buildArrayDataWithUserId(FormRequest $request): array
{
    return [
        ...$request->validated(),
        'user_id' => $request->user()->id
    ];
}

function getUserIdByRequest(Request|FormRequest $req): int
{
    return $req->user()->id;
}
