<?php

namespace App\Http\Requests\Auth;

use Illuminate\Foundation\Http\FormRequest;

class AuthRequest extends FormRequest
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
            'email' => [ 'required', 'min:3', 'max:255', 'email'],
            'password' => [ 'required', 'min:8', 'max:255' ]
        ];
    }

    public function messages()
    {
        return [
            'email.required' => 'O e-mail do usuário é obrigatário.',
            'email.min' => 'O e-mail do usuário deve conter conter no minímo :min caracteres.',
            'email.max' => 'O e-mail do usuário deve conter conter no máximo :max caracteres.',
            'email.email' => 'O e-mail do usuário precisa estar em um formato válido.',
            'password.required' => 'A senha do usuário é obrigatária.',
            'password.min' => 'A senha do usuário deve conter conter no minímo :min caracteres.',
            'password.max' => 'A senha do usuário deve conter conter no máximo :max caracteres.'
        ];
    }
}
