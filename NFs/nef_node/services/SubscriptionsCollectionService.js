/* eslint-disable no-unused-vars */
const Service = require('./Service');

/**
* subscribe to notifications
*
* nefEventExposureSubsc NefEventExposureSubsc 
* returns NefEventExposureSubsc
* */
const createIndividualSubcription = ({ nefEventExposureSubsc }) => new Promise(
  async (resolve, reject) => {
    try {
      resolve(Service.successResponse({
        nefEventExposureSubsc,
      }));
    } catch (e) {
      reject(Service.rejectResponse(
        e.message || 'Invalid input',
        e.status || 405,
      ));
    }
  },
);

module.exports = {
  createIndividualSubcription,
};
