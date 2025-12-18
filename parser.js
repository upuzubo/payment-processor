const fs = require('fs');
const moment = require('moment');
const _ = require('lodash');

class PaymentParser {
  constructor(filePath) {
    this.filePath = filePath;
    this.payments = [];
  }

  readPaymentFile() {
    try {
      const data = fs.readFileSync(this.filePath, 'utf8');
      return data;
    } catch (error) {
      throw new Error(`Failed to read file: ${error.message}`);
    }
  }

  parsePaymentData(data) {
    const lines = data.split('\n');
    lines.forEach((line) => {
      const paymentInfo = line.split(',');
      if (paymentInfo.length === 5) {
        const payment = {
          id: paymentInfo[0],
          amount: parseFloat(paymentInfo[1]),
          currency: paymentInfo[2],
          date: moment(paymentInfo[3], 'YYYY-MM-DD').toDate(),
          status: paymentInfo[4],
        };
        this.payments.push(payment);
      }
    });
  }

  getPayments() {
    return this.payments;
  }

  filterPaymentsByDate(startDate, endDate) {
    return _.filter(this.payments, (payment) => {
      return moment(payment.date).isBetween(startDate, endDate);
    });
  }

  filterPaymentsByStatus(status) {
    return _.filter(this.payments, (payment) => {
      return payment.status === status;
    });
  }
}

module.exports = PaymentParser;