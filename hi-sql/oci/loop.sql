/**
    2018年8月16日 星期四
    循环
**/


begin
  for i in 5..100 loop
    DBMS_OUTPUT.PUT_LINE(i);
    if i = 6 then
      exit; -- 推送循环
    end if;    
  end loop;
  DBMS_OUTPUT.PUT_LINE('循环结束');
end;