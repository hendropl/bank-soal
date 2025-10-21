import ApiHandler from "./api.handler";

// Fungsi untuk mengambil semua soal
export const getAllQuestions = async () => {
  // Asumsi endpoint untuk get all adalah {{baseUrl}}/question
  const res = await ApiHandler.get('/question');
  return res.data;
};

// Fungsi untuk membuat soal
export const createQuestionWithOptions = async (data) => {
  // Menggunakan endpoint {{baseUrl}}/question/options
  const res = await ApiHandler.post('/question/options', data);
  return res.data;
};

// Fungsi untuk menghapus soal berdasarkan ID
export const deleteQuestion = async (id) => {
  // Asumsi endpoint delete adalah {{baseUrl}}/question/:id
  const res = await ApiHandler.delete(`/question/${id}`);
  return res.data;
};

// Anda bisa menambahkan fungsi lain seperti getQuestionById atau updateQuestion di sini nanti