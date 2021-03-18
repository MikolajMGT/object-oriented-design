program sorter;

const n = 101;

procedure Generate(var data: Array of Integer);
var i: Integer;
begin
  Randomize;
  for i := 1 to n do data[i] := Random(n)
end;

procedure Sort(var data: Array of Integer);
var
  i, j, temp: Integer;
begin
  for i := 1 to n do
  begin
    for j := n downTo i do
    begin
      if data[j-1] > data[j] then
      begin
        temp := data[j];
        data[j] := data[j-1];
        data[j-1] := temp;
      end;
    end;
  end;
end;

var
  data: Array[1..n] of Integer;
  i: Integer;

begin
  Generate(data);
  Sort(data);

  for i := 1 to n do Write(data[i], ', ');

  WriteLn()
end.