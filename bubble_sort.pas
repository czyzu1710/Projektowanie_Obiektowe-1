program BubbleSort;

uses sysutils,  math;

var
	rand : Real;
	rand_arr : array[0..99] of Real;
	i : Integer;

procedure generate_random_numbers(n : Integer);
	begin
		for i:= 0 to (n-1) do
			rand_arr[i] := RandomRange(0,100);
	end;

procedure bubble_sort(var arr : array of real);
	var
		tmp : Real;
		i,j : Integer;
	begin
		for i := 0 to Length(arr) - 1 do
			begin
				for j := 0 to Length(arr) - 2 - i do
					begin
						if ( arr[j] > arr[j+1] ) then
							begin
								tmp := arr[j];
								arr[j] := arr[j+1];
								arr[j+1] := tmp;
							end;
					end;
			end;
	end;


begin
	randomize();

	generate_random_numbers(100);
	
	writeln('Random:');
	for i := 0 to Length(rand_arr)-1 do
		writeln(FormatFloat('00', rand_arr[i]));

		

	bubble_sort(rand_arr);
	writeln('Sorted:');
	for i := 0 to Length(rand_arr)-1 do
		writeln(FormatFloat('00', rand_arr[i]));
end.
