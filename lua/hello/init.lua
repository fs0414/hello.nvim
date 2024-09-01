local M = {}

function M.setup()
  vim.api.nvim_create_user_command('Hello', function()
    local result = vim.fn.rpcrequest(0, 'Hello', {})
    print(result)
  end, {})
end

return M
