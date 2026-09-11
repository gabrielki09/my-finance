<?php

use App\Enum\Financial\FinancialAccountTypes;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('financial_accounts', function (Blueprint $table) {
            $table->id();
            $table->uuid('uuid')->unique('un_financial_account_uuid')->index('idx_financial_account_uuid');
            $table->foreignId('user_id')->constrained('users');
            $table->enum('type', array_column(FinancialAccountTypes::cases(), 'value'));
            $table->decimal('initial_balance', 12, 4);
            $table->decimal('current_balance', 12, 4);
            $table->string('name', 255);
            $table->softDeletes();
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('financial_accounts');
    }
};
