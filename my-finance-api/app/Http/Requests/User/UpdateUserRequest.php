<?php

namespace App\Http\Requests\User;

use Illuminate\Contracts\Validation\ValidationRule;
use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Validation\Rule;

class UpdateUserRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'name' => [ 'sometimes', 'min:3', 'max:255'],
            'email' => [ 'sometimes', 'min:3', 'max:255', 'email', Rule::unique('users', 'email')->ignore($this->route('id'))],
            'password' => [ 'sometimes', 'min:8', 'max:255' ]
        ];
    }

    public function messages()
    {
        return [
            'name.min' => 'O nome do usuário deve conter no minímo :min caracteres.',
            'name.max' => 'O nome do usuário deve conter no máximo :max caracteres.',
            'email.min' => 'O e-mail do usuário deve conter conter no minímo :min caracteres.',
            'email.max' => 'O e-mail do usuário deve conter conter no máximo :max caracteres.',
            'email.email' => 'O e-mail do usuário precisa estar em um formato válido.',
            'email.unique' => 'Esse e-mail de usuário já está cadastrado.',
            'password.min' => 'A senha do usuário deve conter conter no minímo :min caracteres.',
            'password.max' => 'A senha do usuário deve conter conter no máximo :max caracteres.'
        ];
    }
}
