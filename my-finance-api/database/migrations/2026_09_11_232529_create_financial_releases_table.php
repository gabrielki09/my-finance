<?php

use App\Enum\Financial\FinancialReleasesStatus;
use App\Enum\Financial\FinancialReleasesType;
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
        Schema::create('financial_releases', function (Blueprint $table) {
            $table->id();
            $table->uuid('uuid')->unique('un_financial_release_uuid')->index('idx_financial_release_uuid');
            $table->foreignId('user_id')->constrained('users');
            $table->foreignUuid('account_id')->constrained('financial_accounts', 'uuid');
            $table->foreignUuid('category_id')->constrained('financial_categories', 'uuid');
            $table->enum('type', array_column(FinancialReleasesType::cases(), 'value'));
            $table->string('description');
            $table->decimal('amount', 12, 2);
            $table->timestamp('transaction_date')->nullable();
            $table->timestamp('due_date')->nullable();
            $table->timestamp('paid_date')->nullable();
            $table->enum('status', array_column(FinancialReleasesStatus::cases(), 'value'))->default(FinancialReleasesStatus::PENDING->value);
            $table->text('notes')->nullable();
            $table->timestamps();
            $table->softDeletes();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('financial_releases');
    }
};
